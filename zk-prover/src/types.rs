use ark_bn254::Fr;
use ark_serialize::{CanonicalDeserialize, CanonicalSerialize};
use serde::{Deserialize, Serialize};

/// Vote proof data
#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct VoteProof {
    /// The Groth16 proof
    pub proof: Vec<u8>,
    /// Public inputs to the circuit
    pub public_inputs: Vec<String>,
}

/// Vote commitment data
#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct VoteCommitment {
    /// The commitment hash
    pub commitment: String,
    /// The choice (kept private until reveal)
    pub choice: u64,
    /// Random nonce (kept private until reveal)
    pub nonce: [u8; 32],
    /// Voter address
    pub voter: String,
}

/// Eligibility proof data
#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct EligibilityProof {
    /// The Groth16 proof
    pub proof: Vec<u8>,
    /// Public inputs (Merkle root)
    pub public_inputs: Vec<String>,
}

/// Merkle tree node
#[derive(Clone, Debug, CanonicalSerialize, CanonicalDeserialize)]
pub struct MerkleNode {
    pub hash: Fr,
}

/// Merkle proof path
#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct MerklePath {
    /// Path from leaf to root
    pub path: Vec<String>,
    /// Indices indicating left/right at each level
    pub indices: Vec<bool>,
}

/// Result type for prover operations
pub type ProverResult<T> = Result<T, crate::error::ProverError>;

/// Configuration for Poseidon hash
#[derive(Clone, Debug)]
pub struct PoseidonConfig {
    /// Number of full rounds
    pub full_rounds: usize,
    /// Number of partial rounds
    pub partial_rounds: usize,
    /// Alpha value
    pub alpha: usize,
}

impl Default for PoseidonConfig {
    fn default() -> Self {
        Self {
            full_rounds: 8,
            partial_rounds: 56,
            alpha: 5,
        }
    }
}
