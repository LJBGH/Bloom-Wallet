import { apiRequest } from './apiClient';

const API_PREFIX = '/api';

function unwrapOk(data) {
  // backend returns model.OkWithData(...), shape likely: { code, msg, data }
  if (data && typeof data === 'object' && 'data' in data) return data.data;
  return data;
}

export async function getAllAccounts() {
  const res = await apiRequest(`${API_PREFIX}/wallet/account/all`);
  const data = unwrapOk(res);
  return (data && data.accounts) || [];
}

export async function getCurrentAccount() {
  const res = await apiRequest(`${API_PREFIX}/wallet/account/current`);
  return unwrapOk(res);
}

export async function switchCurrentAccount(index) {
  const res = await apiRequest(`${API_PREFIX}/wallet/account/switch?index=${encodeURIComponent(index)}`);
  return unwrapOk(res);
}

export async function createMnemonic() {
  const res = await apiRequest(`${API_PREFIX}/wallet/create/mnemonic`);
  return unwrapOk(res);
}

export async function createWallet({ mnemonic } = {}) {
  const q = mnemonic && String(mnemonic).trim().length > 0 ? `?mnemonic=${encodeURIComponent(mnemonic.trim())}` : '';
  const res = await apiRequest(`${API_PREFIX}/wallet/create/wallet${q}`);
  return unwrapOk(res);
}

export async function restoreWallet() {
  const res = await apiRequest(`${API_PREFIX}/wallet/restore`);
  return unwrapOk(res);
}

export async function restoreWithMnemonic(mnemonic) {
  const res = await apiRequest(
    `${API_PREFIX}/wallet/restore/mnemonic?mnemonic=${encodeURIComponent(String(mnemonic || '').trim())}`,
  );
  return unwrapOk(res);
}

export async function restoreWithPrivateKey(privateKey) {
  const res = await apiRequest(
    `${API_PREFIX}/wallet/restore/privateKey?privateKey=${encodeURIComponent(String(privateKey || '').trim())}`,
  );
  return unwrapOk(res);
}
