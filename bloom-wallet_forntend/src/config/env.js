export function getBackendBaseUrl() {
  const injected = import.meta?.env?.VITE_BACKEND_API_URL;
  return injected && injected.trim().length > 0 ? injected.trim() : 'http://localhost:8081';
}
