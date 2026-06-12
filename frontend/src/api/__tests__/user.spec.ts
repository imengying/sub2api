import { beforeEach, describe, expect, it, vi } from 'vitest'

const { post } = vi.hoisted(() => ({
  post: vi.fn(),
}))

vi.mock('@/api/client', () => ({
  apiClient: {
    post,
  },
}))

import { bindEmailIdentity, sendEmailBindingCode } from '@/api/user'

describe('user api account bindings', () => {
  beforeEach(() => {
    post.mockReset()
    post.mockResolvedValue({ data: {} })
  })

  it('sends email binding verification codes', async () => {
    await sendEmailBindingCode('alice@example.com')

    expect(post).toHaveBeenCalledWith('/user/account-bindings/email/send-code', {
      email: 'alice@example.com',
    })
  })

  it('binds or replaces the email identity', async () => {
    const payload = {
      email: 'alice@example.com',
      verify_code: '123456',
      password: 'current-password',
    }

    await bindEmailIdentity(payload)

    expect(post).toHaveBeenCalledWith('/user/account-bindings/email', payload)
  })
})
