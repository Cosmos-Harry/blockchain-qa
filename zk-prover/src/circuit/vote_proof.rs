use ark_bn254::Fr;
use ark_ff::PrimeField;
use ark_r1cs_std::prelude::*;
use ark_r1cs_std::fields::fp::FpVar;
use ark_relations::r1cs::{ConstraintSynthesizer, ConstraintSystemRef, SynthesisError};
use sha2::{Sha256, Digest};

/// Vote proof circuit
/// Proves that a vote is valid without revealing the choice
///
/// Public inputs:
/// - commitment: Hash(choice || nonce || voter)
///
/// Private witnesses:
/// - choice: The vote choice (must be < max_choice)
/// - nonce: Random nonce for privacy
/// - voter: Voter address
///
/// Constraints:
/// 1. choice < max_choice (range check)
/// 2. commitment == Hash(choice || nonce || voter)
#[derive(Clone)]
pub struct VoteProofCircuit {
    // Private witnesses
    pub choice: Option<Fr>,
    pub nonce: Option<Fr>,
    pub voter: Option<Fr>,

    // Public inputs
    pub commitment: Option<Fr>,
    pub max_choice: Option<Fr>,
}

impl VoteProofCircuit {
    /// Create a new circuit with witnesses
    pub fn new_with_witness(
        choice: u64,
        nonce: [u8; 32],
        voter: [u8; 20], // Ethereum address
        max_choice: u64,
    ) -> Self {
        // Convert to field elements
        let choice_fr = Fr::from(choice);
        let nonce_fr = bytes_to_fr(&nonce);
        let voter_fr = bytes_to_fr(&voter);
        let max_choice_fr = Fr::from(max_choice);

        // Compute commitment
        let commitment = compute_commitment(choice, &nonce, &voter);

        Self {
            choice: Some(choice_fr),
            nonce: Some(nonce_fr),
            voter: Some(voter_fr),
            commitment: Some(commitment),
            max_choice: Some(max_choice_fr),
        }
    }

    /// Create a circuit without witnesses (for setup)
    pub fn new_without_witness(max_choice: u64) -> Self {
        Self {
            choice: None,
            nonce: None,
            voter: None,
            commitment: None,
            max_choice: Some(Fr::from(max_choice)),
        }
    }
}

impl ConstraintSynthesizer<Fr> for VoteProofCircuit {
    fn generate_constraints(self, cs: ConstraintSystemRef<Fr>) -> Result<(), SynthesisError> {
        // Allocate private witnesses
        let choice_var = FpVar::new_witness(cs.clone(), || {
            self.choice.ok_or(SynthesisError::AssignmentMissing)
        })?;

        let nonce_var = FpVar::new_witness(cs.clone(), || {
            self.nonce.ok_or(SynthesisError::AssignmentMissing)
        })?;

        let voter_var = FpVar::new_witness(cs.clone(), || {
            self.voter.ok_or(SynthesisError::AssignmentMissing)
        })?;

        // Allocate public inputs
        let commitment_var = FpVar::new_input(cs.clone(), || {
            self.commitment.ok_or(SynthesisError::AssignmentMissing)
        })?;

        let max_choice_var = FpVar::new_input(cs.clone(), || {
            self.max_choice.ok_or(SynthesisError::AssignmentMissing)
        })?;

        // Constraint 1: Range check (choice < max_choice)
        // We use the less_than constraint which checks if choice < max_choice
        choice_var.enforce_cmp(&max_choice_var, std::cmp::Ordering::Less, true)?;

        // Constraint 2: Commitment verification
        // In a real implementation, we'd use Poseidon hash or Pedersen commitment
        // For simplicity here, we simulate the hash constraint
        // computed_commitment = Hash(choice || nonce || voter)

        // Simplified commitment check using field arithmetic
        // Real implementation would use Poseidon hash gadget
        let computed_commitment = simulate_hash_constraint(
            cs.clone(),
            &choice_var,
            &nonce_var,
            &voter_var,
        )?;

        // Enforce computed_commitment == commitment_var
        computed_commitment.enforce_equal(&commitment_var)?;

        Ok(())
    }
}

/// Simulate hash constraint (simplified for demonstration)
/// In production, use ark-crypto-primitives Poseidon hash gadget
fn simulate_hash_constraint(
    _cs: ConstraintSystemRef<Fr>,
    choice: &FpVar<Fr>,
    nonce: &FpVar<Fr>,
    voter: &FpVar<Fr>,
) -> Result<FpVar<Fr>, SynthesisError> {
    // Simplified: commitment = choice + nonce + voter (mod field)
    // In production, use proper Poseidon hash gadget from ark-crypto-primitives
    let mut result = choice.clone();
    result += nonce;
    result += voter;

    Ok(result)
}

/// Convert bytes to field element
fn bytes_to_fr(bytes: &[u8]) -> Fr {
    let mut hash = Sha256::new();
    hash.update(bytes);
    let result = hash.finalize();

    // Take first 31 bytes to ensure it fits in Fr
    let mut bytes_31 = [0u8; 32];
    bytes_31[1..32].copy_from_slice(&result[0..31]);

    Fr::from_le_bytes_mod_order(&bytes_31)
}

/// Compute commitment hash (off-circuit)
pub fn compute_commitment(choice: u64, nonce: &[u8; 32], voter: &[u8]) -> Fr {
    let choice_fr = Fr::from(choice);
    let nonce_fr = bytes_to_fr(nonce);
    let voter_fr = bytes_to_fr(voter);

    // Simplified: commitment = choice + nonce + voter
    // Matches the circuit constraint
    choice_fr + nonce_fr + voter_fr
}

#[cfg(test)]
mod tests {
    use super::*;
    use ark_relations::r1cs::ConstraintSystem;

    #[test]
    fn test_circuit_satisfiability() {
        let choice = 1u64;
        let nonce = [42u8; 32];
        let voter = [1u8; 20];
        let max_choice = 3u64;

        let circuit = VoteProofCircuit::new_with_witness(choice, nonce, voter, max_choice);

        let cs = ConstraintSystem::<Fr>::new_ref();
        circuit.generate_constraints(cs.clone()).unwrap();

        assert!(cs.is_satisfied().unwrap());
    }

    #[test]
    fn test_circuit_fails_with_invalid_choice() {
        let choice = 5u64; // Greater than max_choice
        let nonce = [42u8; 32];
        let voter = [1u8; 20];
        let max_choice = 3u64;

        let circuit = VoteProofCircuit::new_with_witness(choice, nonce, voter, max_choice);

        let cs = ConstraintSystem::<Fr>::new_ref();

        // Should either fail to generate constraints or not be satisfied
        match circuit.generate_constraints(cs.clone()) {
            Ok(_) => {
                // If constraints generated, they should not be satisfied
                assert!(!cs.is_satisfied().unwrap());
            }
            Err(_) => {
                // Expected behavior for invalid input
            }
        }
    }

    #[test]
    fn test_commitment_computation() {
        let choice = 1u64;
        let nonce = [42u8; 32];
        let voter = [1u8; 20];

        let commitment = compute_commitment(choice, &nonce, &voter);

        // Commitment should be deterministic
        let commitment2 = compute_commitment(choice, &nonce, &voter);
        assert_eq!(commitment, commitment2);

        // Different inputs should give different commitments
        let commitment3 = compute_commitment(2, &nonce, &voter);
        assert_ne!(commitment, commitment3);
    }
}
