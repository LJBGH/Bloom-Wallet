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
