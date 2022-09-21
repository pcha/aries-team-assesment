import {Alert, Snackbar} from "@mui/material";

/**
 * Component to show action results, such as a Product creation
 *
 * @param props - {
 *     message: message to show,
 *     success: flag indicating if it's a success or error message,
 *     open: flag to show the result message,
 *     setOpen: setter to change the open flag
 * }
 * @constructor
 */
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