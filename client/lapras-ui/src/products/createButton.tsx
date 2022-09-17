import Button from "@mui/material/Button";
import {AddCircle} from "@mui/icons-material";

function CreateButton(props: { handleShowForm: () => void }) {
    return (<Button
            sx={{position: 'absolute', top: 32, right: 32}}
            onClick={props.handleShowForm} variant="contained" endIcon={<AddCircle/>}>Create</Button>
    )
}

export default CreateButton