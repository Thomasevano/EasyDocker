export interface IAuth {
  auth: (data: IAuthEntity) => Promise<any>
  register: (data: IAuthEntity) => Promise<any>
}

export interface IAuthEntity {
  email: string
  password: string
  passwordConfirm?: string
  remember?: boolean
}

export interface IAuthStatus {
  success: boolean
  errors: IAuthError[]
}

export interface IAuthError {
  path: string
  message: string
}
