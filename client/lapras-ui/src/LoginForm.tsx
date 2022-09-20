import DialogTitle from "@mui/material/DialogTitle";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import {Login, Visibility, VisibilityOff} from "@mui/icons-material";
import {useEffect, useState} from "react";
import {ApiUrl} from "./constants";
import {Alert, Collapse, FormControl, IconButton, InputAdornment, InputLabel, OutlinedInput} from "@mui/material";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import Dialog from "@mui/material/Dialog";

function LoginForm(props: { show: boolean, handleLogIn: () => void, setToken: (token:string) => void }) {
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")
    const [errorMsg, setErrorMsg] = useState("")
    const [showPassword, setShowPassword] = useState(false)

    useEffect(()=>{
        setUsername("")
        setPassword("")
    }, [props.show])

    const login = (username: string, password: string) => {
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
                                props.setToken(body.token)
                            }).then(() => props.handleLogIn())
                        break
                    case 401:
                    case 500:
                    default:
                        res.json().then(body => setErrorMsg(body.error || "unknown error"))
                }
            }).catch(() => setErrorMsg("connection error"))
    }

    const toggleShowPassword = () => {
        setShowPassword(!showPassword)
    }

    const handleMouseDownPassword = () => {}

    return (
        <Dialog open={props.show}>
            <DialogTitle>Log In</DialogTitle>
            <Collapse in={errorMsg != ""}>
                <Alert severity="error">{errorMsg}</Alert>
            </Collapse>
            <DialogContent>
                <TextField
                    autoFocus
                    margin="dense"
                    id="username"
                    label="Username"
                    type="text"
                    fullWidth

                    onChange={e => setUsername(e.target.value)}
                />
                <FormControl
                    variant="outlined"
                    fullWidth
                    sx={{marginTop: 1}}
                >
                    <InputLabel htmlFor="password">Password</InputLabel>
                <OutlinedInput
                    autoFocus
                    margin="dense"
                    id="password"
                    label="Password"
                    type={showPassword ? "text" : "password"}
                    fullWidth
                    onChange={e => setPassword(e.target.value)}
                    endAdornment={<InputAdornment position="end">
                        <IconButton
                            aria-label="toggle password visibility"
                            onClick={toggleShowPassword}
                            onMouseDown={handleMouseDownPassword}
                            edge="end">{showPassword ? <VisibilityOff /> : <Visibility />}</IconButton>
                    </InputAdornment>}
                />
                </FormControl>
            </DialogContent>
            <DialogActions>
                {/*<Button onClick={props.handleClose} variant="outlined" endIcon={<CloseRounded/>}>Cancel</Button>*/}
                <Button onClick={() => {
                    setErrorMsg("")
                    login(username, password)
                }} variant="contained" endIcon={<Login/>}>Log In</Button>
            </DialogActions>
        </Dialog>
    )
}

export default LoginForm