import { request } from '@playwright/test';

export interface Poll {
  id: number;
  contract_address: string;
  question: string;
  options: string[];
  duration: number;
  state: string;
  created_at: string;
  closes_at: string;
}

export interface Vote {
  id: number;
  poll_address: string;
  voter: string;
  commitment: string;
  choice?: number;
  revealed: boolean;
  committed_at: string;
  revealed_at?: string;
}

export interface VoteStats {
  poll_address: string;
  total_votes: number;
  revealed_votes: number;
  pending_reveals: number;
}

export interface Results {
  poll_address: string;
  vote_counts: number[];
  total_votes: number;
  tallied_at: string;
}

export class APIHelper {
  baseURL: string;

  constructor(baseURL: string) {
    this.baseURL = baseURL;
  }

  async getPoll(pollAddress: string): Promise<Poll | null> {
    const context = await request.newContext({ baseURL: this.baseURL });
    try {
      const response = await context.get(`/api/polls/${pollAddress}`);
      if (response.status() === 404) {
        return null;
      }
      if (!response.ok()) {
        throw new Error(`API error: ${response.status()} ${await response.text()}`);
      }
      return await response.json();
    } finally {
      await context.dispose();
    }
  }

  async listPolls(state?: string, limit: number = 20, offset: number = 0): Promise<Poll[]> {
    const context = await request.newContext({ baseURL: this.baseURL });
    try {
      const params = new URLSearchParams({
        limit: limit.toString(),
        offset: offset.toString(),
      });
      if (state) {
        params.append('state', state);
      }

      const response = await context.get(`/api/polls?${params}`);
      if (!response.ok()) {
        throw new Error(`API error: ${response.status()} ${await response.text()}`);
      }
      const data = await response.json();
      return data.polls || [];
    } finally {
      await context.dispose();
    }
  }

  async getVotes(pollAddress: string, revealedOnly: boolean = false): Promise<Vote[]> {
    const context = await request.newContext({ baseURL: this.baseURL });
    try {
      const params = new URLSearchParams({
        revealed_only: revealedOnly.toString(),
      });

      const response = await context.get(`/api/polls/${pollAddress}/votes?${params}`);
      if (!response.ok()) {
        throw new Error(`API error: ${response.status()} ${await response.text()}`);
      }
      const data = await response.json();
      return data.votes || [];
    } finally {
      await context.dispose();
    }
  }

  async getVoteStats(pollAddress: string): Promise<VoteStats> {
    const context = await request.newContext({ baseURL: this.baseURL });
    try {
      const response = await context.get(`/api/polls/${pollAddress}/stats`);
      if (!response.ok()) {
        throw new Error(`API error: ${response.status()} ${await response.text()}`);
      }
      return await response.json();
    } finally {
      await context.dispose();
    }
  }

  async getResults(pollAddress: string): Promise<Results | null> {
    const context = await request.newContext({ baseURL: this.baseURL });
    try {
      const response = await context.get(`/api/polls/${pollAddress}/results`);
      if (response.status() === 404) {
        return null;
      }
      if (!response.ok()) {
        throw new Error(`API error: ${response.status()} ${await response.text()}`);
      }
      return await response.json();
    } finally {
      await context.dispose();
    }
  }

  async healthCheck(): Promise<boolean> {
    const context = await request.newContext({ baseURL: this.baseURL });
    try {
      const response = await context.get('/health');
      return response.ok();
    } catch {
      return false;
    } finally {
      await context.dispose();
    }
  }

  async waitForIndexer(pollAddress: string, timeoutMs: number = 30000): Promise<Poll> {
    const startTime = Date.now();
    while (Date.now() - startTime < timeoutMs) {
      const poll = await this.getPoll(pollAddress);
      if (poll) {
        return poll;
      }
      await new Promise(resolve => setTimeout(resolve, 1000)); // Wait 1 second
    }
    throw new Error(`Indexer did not index poll ${pollAddress} within ${timeoutMs}ms`);
  }
}
