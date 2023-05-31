import {useState} from "react";
import CreateForm from "./CreateForm";
import CreateButton from "./CreateButton";

/**
 * Component with the product creation related components
 * @param props - {
 *     handleCreate: handler to be called in order to create the product
 * }
 * @constructor
 */
function Creator(props: { handleCreate: (name: string, description: string) => void }) {
    const [open, setOpen] = useState(false)

    const handleOpen = () => {
        setOpen(true)
    }

    const handleClose = () => {
        setOpen(false)
    }

    const handleCreate = (name: string, description: string) => {
        props.handleCreate(name, description)
        setOpen(false)
    }

    return (
        <div>
            <CreateButton handleShowForm={handleOpen}/>
            <CreateForm open={open} handleClose={handleClose} handleCreate={handleCreate}/>
        </div>
    )
}

export default Creator