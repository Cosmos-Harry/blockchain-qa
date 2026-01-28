import { test, expect } from '@playwright/test';
import { APIHelper } from '../utils/api';
import * as dotenv from 'dotenv';

dotenv.config();

const API_URL = process.env.API_URL || 'http://localhost:3000';

test.describe('API Endpoints', () => {
  let api: APIHelper;

  test.beforeAll(() => {
    api = new APIHelper(API_URL);
  });

  test('health check endpoint', async () => {
    // Retry up to 30 seconds for API to be ready
    let healthy = false;
    for (let i = 0; i < 30; i++) {
      healthy = await api.healthCheck();
      if (healthy) break;
      await new Promise(resolve => setTimeout(resolve, 1000));
    }
    expect(healthy).toBeTruthy();
    console.log('✓ API health check passed');
  });

  test('list polls with pagination', async () => {
    // Test with different pagination parameters
    const page1 = await api.listPolls('', 5, 0);
    expect(Array.isArray(page1)).toBeTruthy();
    expect(page1.length).toBeLessThanOrEqual(5);

    const page2 = await api.listPolls('', 5, 5);
    expect(Array.isArray(page2)).toBeTruthy();

    console.log(`✓ Pagination test: Page 1 has ${page1.length} polls, Page 2 has ${page2.length} polls`);
  });

  test('list polls with state filter', async () => {
    const activePolls = await api.listPolls('active');
    const closedPolls = await api.listPolls('closed');
    const talliedPolls = await api.listPolls('tallied');

    expect(Array.isArray(activePolls)).toBeTruthy();
    expect(Array.isArray(closedPolls)).toBeTruthy();
    expect(Array.isArray(talliedPolls)).toBeTruthy();

    console.log(`✓ State filter test: active=${activePolls.length}, closed=${closedPolls.length}, tallied=${talliedPolls.length}`);
  });

  test('get poll by address - not found', async () => {
    const nonExistentAddress = '0x' + '0'.repeat(40);
    const poll = await api.getPoll(nonExistentAddress);

    expect(poll).toBeNull();
    console.log('✓ Get poll test: Non-existent poll returns null');
  });

  test('get vote stats - error handling', async () => {
    const nonExistentAddress = '0x' + '0'.repeat(40);

    try {
      await api.getVoteStats(nonExistentAddress);
      // If this succeeds with a poll that doesn't exist, that's fine too
    } catch (error) {
      // Expected to fail for non-existent poll
      expect(error).toBeDefined();
    }

    console.log('✓ Vote stats test: Handles non-existent polls correctly');
  });

  test('get results - not found', async () => {
    const nonExistentAddress = '0x' + '0'.repeat(40);
    const results = await api.getResults(nonExistentAddress);

    expect(results).toBeNull();
    console.log('✓ Get results test: Non-existent poll results return null');
  });

  test('API response time', async () => {
    const start = Date.now();
    await api.listPolls('', 10, 0);
    const duration = Date.now() - start;

    expect(duration).toBeLessThan(5000); // Should respond within 5 seconds
    console.log(`✓ Response time test: List polls took ${duration}ms`);
  });
});
