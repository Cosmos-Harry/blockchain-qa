use ark_bn254::{Bn254, Fr};
use ark_groth16::{Proof, ProvingKey, VerifyingKey};
use ark_serialize::{CanonicalDeserialize, CanonicalSerialize};
use ark_std::rand::rngs::StdRng;
use ark_std::rand::SeedableRng;

use crate::circuit::VoteProofCircuit;
use crate::error::ProverError;
use crate::types::{ProverResult, VoteProof};

/// Main prover struct for generating and verifying ZK proofs
pub struct Prover {
    proving_key: ProvingKey<Bn254>,
    verifying_key: VerifyingKey<Bn254>,
}

impl Prover {
    /// Perform Groth16 trusted setup
    /// In production, use a multi-party computation ceremony
    pub fn setup(max_choice: u64) -> ProverResult<Self> {
        let mut rng = StdRng::from_entropy();

        // Create circuit without witnesses for setup
        let circuit = VoteProofCircuit::new_without_witness(max_choice);

        // Generate proving key (contains verifying key)
        let proving_key = ark_groth16::Groth16::<Bn254>::generate_random_parameters_with_reduction(
            circuit, &mut rng,
        )
        .map_err(|e| ProverError::SynthesisError(e.to_string()))?;

        let verifying_key = proving_key.vk.clone();

        Ok(Self {
            proving_key,
            verifying_key,
        })
    }

    /// Load keys from bytes
    pub fn from_keys(pk_bytes: &[u8], vk_bytes: &[u8]) -> ProverResult<Self> {
        let proving_key = ProvingKey::deserialize_compressed(pk_bytes)?;
        let verifying_key = VerifyingKey::deserialize_compressed(vk_bytes)?;

        Ok(Self {
            proving_key,
            verifying_key,
        })
    }

    /// Generate a vote proof
    pub fn prove_vote(
        &self,
        choice: u64,
        nonce: [u8; 32],
        voter: [u8; 20],
        max_choice: u64,
    ) -> ProverResult<VoteProof> {
        let mut rng = StdRng::from_entropy();

        // Create circuit with witnesses
        let circuit = VoteProofCircuit::new_with_witness(choice, nonce, voter, max_choice);

        // Compute commitment (public input)
        let commitment = crate::circuit::vote_proof::compute_commitment(choice, &nonce, &voter);
        let max_choice_fr = Fr::from(max_choice);

        // Generate proof
        let proof = ark_groth16::Groth16::<Bn254>::create_random_proof_with_reduction(
            circuit,
            &self.proving_key,
            &mut rng,
        )
        .map_err(|e| ProverError::ProofGenerationFailed(e.to_string()))?;

        // Serialize proof
        let mut proof_bytes = Vec::new();
        proof
            .serialize_compressed(&mut proof_bytes)
            .map_err(|e| ProverError::SerializationError(e.to_string()))?;

        // Serialize public inputs (commitment, max_choice)
        let public_inputs = vec![
            fr_to_hex_string(&commitment),
            fr_to_hex_string(&max_choice_fr),
        ];

        Ok(VoteProof {
            proof: proof_bytes,
            public_inputs,
        })
    }

    /// Verify a vote proof
    pub fn verify_vote(&self, proof: &VoteProof) -> ProverResult<bool> {
        // Deserialize proof
        let proof_obj = Proof::deserialize_compressed(&proof.proof[..])
            .map_err(|e| ProverError::SerializationError(e.to_string()))?;

        // Parse public inputs (commitment, max_choice)
        if proof.public_inputs.len() < 2 {
            return Err(ProverError::InvalidParameters(
                "Expected 2 public inputs (commitment, max_choice)".to_string(),
            ));
        }

        let commitment = hex_string_to_fr(&proof.public_inputs[0])?;
        let max_choice = hex_string_to_fr(&proof.public_inputs[1])?;

        let public_inputs = vec![commitment, max_choice];

        // Prepare verifying key for efficient verification
        let prepared_vk = ark_groth16::prepare_verifying_key(&self.verifying_key);

        // Verify proof
        let result =
            ark_groth16::Groth16::<Bn254>::verify_proof(&prepared_vk, &proof_obj, &public_inputs)
                .map_err(|_| ProverError::VerificationFailed)?;

        Ok(result)
    }

    /// Export proving key as bytes
    pub fn export_proving_key(&self) -> ProverResult<Vec<u8>> {
        let mut bytes = Vec::new();
        self.proving_key
            .serialize_compressed(&mut bytes)
            .map_err(|e| ProverError::SerializationError(e.to_string()))?;
        Ok(bytes)
    }

    /// Export verifying key as bytes
    pub fn export_verifying_key(&self) -> ProverResult<Vec<u8>> {
        let mut bytes = Vec::new();
        self.verifying_key
            .serialize_compressed(&mut bytes)
            .map_err(|e| ProverError::SerializationError(e.to_string()))?;
        Ok(bytes)
    }

    /// Get verifying key reference
    pub fn verifying_key(&self) -> &VerifyingKey<Bn254> {
        &self.verifying_key
    }
}

/// Convert field element to hex string
fn fr_to_hex_string(fr: &Fr) -> String {
    let mut bytes = Vec::new();
    fr.serialize_compressed(&mut bytes).unwrap();
    hex::encode(bytes)
}

/// Convert hex string to field element
fn hex_string_to_fr(hex: &str) -> ProverResult<Fr> {
    let bytes = hex::decode(hex).map_err(|e| ProverError::SerializationError(e.to_string()))?;

    Fr::deserialize_compressed(&bytes[..])
        .map_err(|e| ProverError::SerializationError(e.to_string()))
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_setup() {
        let prover = Prover::setup(3).unwrap();
        assert!(prover.export_proving_key().is_ok());
        assert!(prover.export_verifying_key().is_ok());
    }

    #[test]
    fn test_prove_and_verify() {
        let prover = Prover::setup(3).unwrap();

        let choice = 1u64;
        let nonce = [42u8; 32];
        let voter = [1u8; 20];

        // Generate proof
        let proof = prover.prove_vote(choice, nonce, voter, 3).unwrap();

        // Verify proof
        let result = prover.verify_vote(&proof).unwrap();
        assert!(result);
    }

    #[test]
    fn test_invalid_proof_fails() {
        let prover = Prover::setup(3).unwrap();

        // Generate valid proof
        let proof = prover.prove_vote(1, [42u8; 32], [1u8; 20], 3).unwrap();

        // Tamper with proof
        let mut bad_proof = proof.clone();
        bad_proof.proof[0] ^= 0xFF;

        // Verification should fail
        let result = prover.verify_vote(&bad_proof);
        assert!(result.is_err() || !result.unwrap());
    }
}
