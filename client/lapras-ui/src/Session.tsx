import {useCookies} from "react-cookie";

// type session = {
//     token: () => string,
//     username: () => string,
//     setToken: (token: string) => void
//     setUsername: (username: string) => void
//     clear: () => void
// }

const Session = function () {
    const [cookies, setCookies, removeCookies] = useCookies(['session_token', 'session_username'])

    return {
        token: () => cookies.session_token.toString(),
        username: () => cookies.session_username.toString(),
        setToken: (tkn: string) => setCookies('session_token', tkn),
        setUsername: (username: string) => setCookies('session_username', username),
        clear: () => {
            removeCookies("session_token")
            removeCookies("session_username")
        }
    }
}()

export default Session