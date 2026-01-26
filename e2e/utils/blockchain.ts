import { ethers } from 'ethers';

export interface TestAccount {
  address: string;
  privateKey: string;
  signer: ethers.Wallet;
}

export class BlockchainHelper {
  provider: ethers.JsonRpcProvider;
  accounts: TestAccount[];

  constructor(rpcUrl: string, privateKeys: string[]) {
    this.provider = new ethers.JsonRpcProvider(rpcUrl);
    this.accounts = privateKeys.map(pk => {
      const wallet = new ethers.Wallet(pk, this.provider);
      return {
        address: wallet.address,
        privateKey: pk,
        signer: wallet,
      };
    });
  }

  async getBlockNumber(): Promise<number> {
    return await this.provider.getBlockNumber();
  }

  async mineBlocks(count: number): Promise<void> {
    for (let i = 0; i < count; i++) {
      await this.provider.send('evm_mine', []);
    }
  }

  async increaseTime(seconds: number): Promise<void> {
    await this.provider.send('evm_increaseTime', [seconds]);
    await this.provider.send('evm_mine', []);
  }

  async snapshot(): Promise<string> {
    return await this.provider.send('evm_snapshot', []);
  }

  async revert(snapshotId: string): Promise<void> {
    await this.provider.send('evm_revert', [snapshotId]);
  }

  async getBalance(address: string): Promise<bigint> {
    return await this.provider.getBalance(address);
  }

  async waitForTransaction(txHash: string, confirmations: number = 1): Promise<any> {
    return await this.provider.waitForTransaction(txHash, confirmations);
  }
}

export function computeCommitment(choice: number, nonce: string, voter: string): string {
  // This should match the contract/circuit commitment computation
  // For now, using keccak256 as placeholder
  const data = ethers.solidityPacked(
    ['uint256', 'bytes32', 'address'],
    [choice, nonce, voter]
  );
  return ethers.keccak256(data);
}

export function generateNonce(): string {
  return ethers.hexlify(ethers.randomBytes(32));
}

export function generateMerkleRoot(voters: string[]): string {
  // Simplified Merkle root generation
  // In production, this would properly compute the Merkle tree
  if (voters.length === 0) {
    return ethers.ZeroHash;
  }

  // Hash all voters together as a simple root
  const concatenated = voters.join('');
  return ethers.keccak256(ethers.toUtf8Bytes(concatenated));
}
