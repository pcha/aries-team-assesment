import {Add} from "@mui/icons-material";
import {Fab} from "@mui/material";

function CreateButton(props: { handleShowForm: () => void }) {
    return (<Fab variant="extended"
                 color="primary"
                 sx={{
                     position: 'fixed',
                     bottom: 32,
                     right: 32,
                 }}
            onClick={props.handleShowForm}>Create <Add /></Fab>
    )
}

export default CreateButton