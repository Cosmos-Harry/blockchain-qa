use thiserror::Error;

/// Errors that can occur during proof generation and verification
#[derive(Error, Debug)]
pub enum ProverError {
    #[error("Circuit synthesis error: {0}")]
    SynthesisError(String),

    #[error("Proof generation failed: {0}")]
    ProofGenerationFailed(String),

    #[error("Proof verification failed")]
    VerificationFailed,

    #[error("Invalid parameters: {0}")]
    InvalidParameters(String),

    #[error("Serialization error: {0}")]
    SerializationError(String),

    #[error("Invalid choice: {0} >= max {1}")]
    InvalidChoice(u64, u64),

    #[error("Invalid commitment")]
    InvalidCommitment,

    #[error("Invalid Merkle proof")]
    InvalidMerkleProof,

    #[error("IO error: {0}")]
    IoError(#[from] std::io::Error),

    #[error("Constraint system error: {0}")]
    ConstraintSystemError(String),
}

impl From<ark_relations::r1cs::SynthesisError> for ProverError {
    fn from(err: ark_relations::r1cs::SynthesisError) -> Self {
        ProverError::SynthesisError(err.to_string())
    }
}

impl From<ark_serialize::SerializationError> for ProverError {
    fn from(err: ark_serialize::SerializationError) -> Self {
        ProverError::SerializationError(err.to_string())
    }
}
