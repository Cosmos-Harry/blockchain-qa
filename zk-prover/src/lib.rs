pub mod circuit;
pub mod error;
pub mod prover;
pub mod types;

pub use circuit::VoteProofCircuit;
pub use error::ProverError;
pub use prover::Prover;
pub use types::{ProverResult, VoteCommitment, VoteProof};

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_basic_imports() {
        // Ensure all modules are accessible
        let _ = ProverError::InvalidCommitment;
    }
}