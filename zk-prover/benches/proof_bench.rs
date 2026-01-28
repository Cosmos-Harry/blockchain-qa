use criterion::{black_box, criterion_group, criterion_main, Criterion};
use zk_prover::Prover;

fn bench_proof_generation(c: &mut Criterion) {
    let prover = Prover::setup(3).unwrap();

    c.bench_function("vote proof generation", |b| {
        b.iter(|| {
            let choice = black_box(1u64);
            let nonce = black_box([42u8; 32]);
            let voter = black_box([1u8; 20]);

            prover.prove_vote(choice, nonce, voter, 3).unwrap()
        });
    });
}

fn bench_proof_verification(c: &mut Criterion) {
    let prover = Prover::setup(3).unwrap();
    let proof = prover.prove_vote(1, [42u8; 32], [1u8; 20], 3).unwrap();

    c.bench_function("vote proof verification", |b| {
        b.iter(|| prover.verify_vote(black_box(&proof)).unwrap());
    });
}

criterion_group!(benches, bench_proof_generation, bench_proof_verification);
criterion_main!(benches);
