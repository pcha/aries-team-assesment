/**
 * App name to show in the AppBar
 */
export const AppName = "LAPRAS"

/**
 * Api base url
 */
export const ApiUrl:string = import.meta.env.VITE_API_URL

/**
 * Minutes that it should wait for renew the token
 */
export const MinutesToRenewToken = import.meta.env.VITE_MINUTES_TO_RENEW_TOKEN || 13
