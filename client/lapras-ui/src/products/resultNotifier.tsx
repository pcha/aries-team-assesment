import {Alert, Snackbar} from "@mui/material";

function ResultNotifier(props: {message:string, success:boolean, open: boolean, setOpen: (a: boolean) => void}) {
    const handleClose = () => props.setOpen(false)
    return <Snackbar
        open={props.open}
        autoHideDuration={3000}
        onClose={handleClose}
    >
        <Alert onClose={handleClose} severity={props.success? "success": "error"} sx={{ width: '100%'}}>
            {props.message}
        </Alert>
    </Snackbar>
}

export default ResultNotifier