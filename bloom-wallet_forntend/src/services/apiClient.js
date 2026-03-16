import { getBackendBaseUrl } from '../config/env';

function normalizeBaseUrl(url) {
  return url.replace(/\/+$/, '');
}

export function getAuthToken() {
  try {
    return localStorage.getItem('bloom_wallet_token') || '';
  } catch {
    return '';
  }
}

export function setAuthToken(token) {
  try {
    if (!token) localStorage.removeItem('bloom_wallet_token');
    else localStorage.setItem('bloom_wallet_token', token);
  } catch {
    // ignore
  }
}

export class ApiError extends Error {
  constructor(message, { status, data } = {}) {
    super(message);
    this.name = 'ApiError';
    this.status = status;
    this.data = data;
  }
}

export async function apiRequest(path, { method = 'GET', headers, body } = {}) {
  const baseUrl = normalizeBaseUrl(getBackendBaseUrl());
  const url = path.startsWith('http') ? path : `${baseUrl}${path.startsWith('/') ? '' : '/'}${path}`;

  const token = getAuthToken();
  const finalHeaders = {
    Accept: 'application/json',
    ...(body != null ? { 'Content-Type': 'application/json' } : {}),
    ...(token ? { Authorization: token.startsWith('Bearer ') ? token : `Bearer ${token}` } : {}),
    ...(headers || {}),
  };

  const res = await fetch(url, {
    method,
    headers: finalHeaders,
    body: body == null ? undefined : JSON.stringify(body),
  });

  const text = await res.text();
  const data = text ? safeJsonParse(text) : null;

  if (!res.ok) {
    const msg = (data && (data.message || data.error)) || `Request failed (${res.status})`;
    throw new ApiError(msg, { status: res.status, data });
  }
  return data;
}

function safeJsonParse(s) {
  try {
    return JSON.parse(s);
  } catch {
    return s;
  }
}
