import {useCookies} from "react-cookie";
import {useEffect, useState} from "react";
import {ApiUrl} from "../etc/constants";

// Contains session related states and function
export type Session = {
    token: string,
    username: string,
    isLoggedIn: boolean,
    logIn: (username: string, password: string) => void,
    logOut: () => void,
    loginResultMessage: string
}

// Hook to handle user session
function useSession(): Session {
    const [firstRender, setFirstRender] = useState(true)
    const [cookies, setCookie, removeCookie] = useCookies(['token', 'username'])
    const [loggedIn, setLoggedIn] = useState(cookies.token != "") // If there is a previous token in the cookies it renders as logged
    const [loginResultMessage, setLoginResultMessage] = useState("")

    // state representing the cookies values. This is needed to expose these values typed as 'string'. 
    //Otherwise, typing these values will prevent them to get updated.
    const [username, setUsernameState] = useState<string>(cookies.username)
    const [token, setTokenState] =useState<string>(cookies.token)
    // use effect listening cookies changes and updating the corresponding states
    useEffect(() => setUsernameState(cookies.username), [cookies.username])
    useEffect(() => setTokenState(cookies.token), [cookies.token])
    // use effect to handle loggedIn. It depends on the token state because it is this State the one that provides the token values to the different components.
    // Otherwise, it could result in an eventual 401 response.
    useEffect(() => setLoggedIn(!!token), [token])


    // wrapper to set the token in the cookies
    const setToken = (tkn: string) => setCookie('token', tkn);
    // wrapper to set the username in the cookies
    const setUsername = (username: string) => setCookie('username', username)

    const logIn = (username: string, password: string) => {
        setLoginResultMessage("")
        fetch(ApiUrl + "/users/login", {
            method: "POST",
            body: JSON.stringify({
                username: username,
                password: password
            })
        })
            .then(res => {
                switch (res.status) {
                    case 200:
                        res.json()
                            .then(body => {
                                setToken(body.token)
                            })
                        break
                    case 401:
                    case 500:
                    default:
                        res.json().then(body => setLoginResultMessage(body.error || "unknown error"))
                }
            }).catch(() => setLoginResultMessage("Connection error"))
    }

    const logOut = () => {
        removeCookie('token')
        removeCookie('username')
        setLoggedIn(false)
    }

    // Calls the renovation api endpoint and sets the new token, if it fails it will log out
    const renewToken = () => {
        if (!loggedIn || !cookies.token) {
            logOut()
            return
        }
        fetch(ApiUrl + "/users/token/renew", {
            method: 'POST',
            headers: {
                'Authorization': "Bearer " + cookies.token
            }
        }).then(res => {
            switch (res.status) {
                case 200:
                    res.json()
                        .then(body => {
                            setToken(body.token)
                            setUsername(body.claims.username)
                        })
                    break
                case 401:
                case 500:
                default:
                    logOut()
            }
        })
    }

    // To be executed only on the first render
    useEffect(() => {
        // mark the first render as executed
        setFirstRender(false)
        // If there is a token on a first render I do not know when it will expire, so it forces a renovation
        renewToken()
    }, [])

    // When the token changes, it schedules a renovation before expiration
    useEffect(() => {
        // In the first render the token in renewed to avoid expiration, so the renovation will be already scheduled.
        // Also, if the token is empty, it does not make sense to schedule a renovation
        if (!firstRender && cookies.token) {
            setTimeout(renewToken, 13 * 60 * 1000) // renew token in 13 minutes
        }
    }, [cookies.token])


    return {
        token: token,
        username: username,
        isLoggedIn: loggedIn,
        logIn: logIn, //executes log In
        logOut: logOut, //executes log out
        loginResultMessage: loginResultMessage
    }
}

export default useSession