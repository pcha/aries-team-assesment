import './App.css'
import Products from "./products";
import {Container} from "@mui/material";
import Header from "./header";
import {useState} from "react";
import LoginForm from "./LoginForm";
import {useCookies} from "react-cookie";

function App() {
    const [loggedIn, setLoggedIn] = useState(false)
    const [cookies, setCookies, removeCookies] = useCookies(['token'])
    const [filterTerm, setFilterTerm] = useState("")

    const handleLoggedIn = () => {
        setLoggedIn(true)
    }
    const logOut = () => {
        removeCookies('token')
        setLoggedIn(false)
    }
    const setToken = (token:string) => setCookies('token', token);

    return (
        <div className="App" style={{overflow: "hidden"}}>
            <Header loggedIn={loggedIn} handleLogOut={() => setLoggedIn(false)} handleSearch={setFilterTerm}/>
            <LoginForm show={!loggedIn} handleLogIn={handleLoggedIn} setToken={setToken}/>
            {loggedIn ? <Container id="content" sx={{}}>
                <Products apiURL={import.meta.env.API_URL} apiToken={cookies.token.toString()} logOut={logOut} filterTerm={filterTerm}/>
            </Container>: ""}
        </div>
    );
}

export default App
