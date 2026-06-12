const OAUTH_AFFILIATE_CODE_KEY = 'oauth_aff_code'
const AFFILIATE_REFERRAL_CODE_KEY = 'affiliate_referral_code'

export function normalizeOAuthAffiliateCode(_value?: unknown): string {
  return ''
}

export function pickOAuthAffiliateCode(..._values: unknown[]): string {
  return ''
}

export function storeAffiliateReferralCode(_value?: unknown, _now = Date.now()): void {
  clearAffiliateReferralCode()
}

export function loadAffiliateReferralCode(_now = Date.now()): string {
  clearAffiliateReferralCode()
  return ''
}

export function clearAffiliateReferralCode(): void {
  if (typeof window === 'undefined') {
    return
  }
  try {
    window.localStorage.removeItem(AFFILIATE_REFERRAL_CODE_KEY)
  } catch {
    // Ignore browser storage failures.
  }
}

export function resolveAffiliateReferralCode(..._values: unknown[]): string {
  clearAffiliateReferralCode()
  return ''
}

export function storeOAuthAffiliateCode(_value?: unknown): void {
  clearOAuthAffiliateCode()
}

export function loadOAuthAffiliateCode(): string {
  clearOAuthAffiliateCode()
  return ''
}

export function clearOAuthAffiliateCode(): void {
  if (typeof window === 'undefined') {
    return
  }
  try {
    window.sessionStorage.removeItem(OAUTH_AFFILIATE_CODE_KEY)
  } catch {
    // Ignore browser storage failures.
  }
}

export function clearAllAffiliateReferralCodes(): void {
  clearOAuthAffiliateCode()
  clearAffiliateReferralCode()
}

export function oauthAffiliatePayload(_value?: unknown): Record<string, never> {
  return {}
}
