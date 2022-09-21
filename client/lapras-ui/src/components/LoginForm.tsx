import DialogTitle from "@mui/material/DialogTitle";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import {Login, Visibility, VisibilityOff} from "@mui/icons-material";
import {useEffect, useState} from "react";
import {Alert, Collapse, FormControl, IconButton, InputAdornment, InputLabel, OutlinedInput} from "@mui/material";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import Dialog from "@mui/material/Dialog";
import {Session} from "../hooks/useSession";

function LoginForm(props: { session: Session }) {
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")
    const [showPassword, setShowPassword] = useState(false)

    useEffect(() => {
        setUsername("")
        setPassword("")
    }, [props.session.isLoggedIn])

    const toggleShowPassword = () => {
        setShowPassword(!showPassword)
    }

    return (
        <Dialog open={!props.session.isLoggedIn}>
            <DialogTitle>Log In</DialogTitle>
            <Collapse in={props.session.loginResultMessage != ""}>
                <Alert severity="error">{props.session.loginResultMessage}</Alert>
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
                                edge="end">{showPassword ? <VisibilityOff/> : <Visibility/>}</IconButton>
                        </InputAdornment>}
                    />
                </FormControl>
            </DialogContent>
            <DialogActions>
                <Button onClick={() => {
                    props.session.logIn(username, password)
                }} variant="contained" endIcon={<Login/>}>Log In</Button>
            </DialogActions>
        </Dialog>
    )
}

export default LoginForm