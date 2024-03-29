import {useEffect, useState} from "react";
import Button from "@mui/material/Button";
import {CloseRounded, Send} from "@mui/icons-material";
import Dialog from "@mui/material/Dialog";
import DialogTitle from "@mui/material/DialogTitle";
import DialogContent from "@mui/material/DialogContent";
import TextField from "@mui/material/TextField";
import DialogActions from "@mui/material/DialogActions";

/**
 * Component With the form to create a new Product
 * @param props - {
 *     open: flag indicating to show or not the Form
 *     handleClose: handler to call when the form should be closed
 *     handleCreate: handler to call when the submit button is closed
 * }
 * @constructor
 */
function CreateForm(props: { open: boolean, handleClose: () => void, handleCreate: (name: string, description: string) => void }) {
    const [name, setName] = useState("")
    const [description, setDescription] = useState("")

    useEffect(() => {
        setName("")
        setDescription("")
    }, [props.open])
    return (
        <Dialog open={props.open} onClose={props.handleClose} fullWidth maxWidth="lg">
            <DialogTitle>Create Product</DialogTitle>
            <DialogContent>
                <TextField
                    autoFocus
                    margin="dense"
                    id="name"
                    label="Name"
                    type="text"
                    fullWidth
                    required={true}
                    onChange={e => setName(e.target.value)}
                />
                <TextField
                    autoFocus
                    margin="dense"
                    id="description"
                    label="Description"
                    type="text"
                    fullWidth
                    required={true}
                    onChange={e => setDescription(e.target.value)}
                    multiline
                    rows={8}
                />
            </DialogContent>
            <DialogActions>
                <Button onClick={props.handleClose} variant="outlined" endIcon={<CloseRounded/>}>Cancel</Button>
                <Button onClick={() => {props.handleCreate(name, description)}} variant="contained" endIcon={<Send/>}>Create</Button>
            </DialogActions>
        </Dialog>

    )
}

export default CreateForm