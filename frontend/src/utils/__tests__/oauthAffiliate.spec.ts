import { beforeEach, describe, expect, it, vi } from 'vitest'
import {
  clearAffiliateReferralCode,
  clearOAuthAffiliateCode,
  loadAffiliateReferralCode,
  loadOAuthAffiliateCode,
  oauthAffiliatePayload,
  resolveAffiliateReferralCode,
  storeAffiliateReferralCode,
  storeOAuthAffiliateCode
} from '@/utils/oauthAffiliate'

describe('oauthAffiliate', () => {
  beforeEach(() => {
    localStorage.clear()
    sessionStorage.clear()
    vi.useRealTimers()
  })

  it('ignores affiliate referral codes', () => {
    expect(resolveAffiliateReferralCode(' 5579J7CFG9PF ')).toBe('')
    expect(loadAffiliateReferralCode()).toBe('')
    expect(localStorage.getItem('affiliate_referral_code')).toBeNull()
  })

  it('clears stored affiliate referral code when touched', () => {
    const now = Date.UTC(2026, 0, 1)
    storeAffiliateReferralCode('AFF123', now)

    expect(loadAffiliateReferralCode(now)).toBe('')
    expect(localStorage.getItem('affiliate_referral_code')).toBeNull()
  })

  it('clears oauth transient code instead of storing it', () => {
    storeAffiliateReferralCode('PERSISTED')
    storeOAuthAffiliateCode('OAUTH')

    expect(loadAffiliateReferralCode()).toBe('')
    expect(loadOAuthAffiliateCode()).toBe('')

    clearOAuthAffiliateCode()
    expect(loadOAuthAffiliateCode()).toBe('')
    expect(loadAffiliateReferralCode()).toBe('')

    clearAffiliateReferralCode()
    expect(loadAffiliateReferralCode()).toBe('')
  })

  it('does not build affiliate payload fields', () => {
    expect(oauthAffiliatePayload('AFF123')).toEqual({})
  })
})
