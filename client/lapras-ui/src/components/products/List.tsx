import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import {Alert} from "@mui/material";
import {Product} from "../../hooks/useProducts";

/**
 * Component with a table listing the given Products
 *
 * @param props - {
 *     products: products to list
 *     maxHeight: max component height
 * }
 * @constructor
 */
function ProductsList(props:{products:Product[], maxHeight: number}) {
    const headerColor = ""

    return (props.products.length ? <TableContainer component={Paper} sx={{maxHeight: props.maxHeight}}>
        <Table stickyHeader>
            <TableHead sx={{backgroundColor: "black"}}>
                <TableRow>
                    <TableCell sx={{backgroundColor: headerColor}}>ID</TableCell>
                    <TableCell sx={{backgroundColor: headerColor}}>Name</TableCell>
                    <TableCell sx={{backgroundColor: headerColor}}>Description</TableCell>
                </TableRow>
            </TableHead>
            <TableBody>
                {props.products.map((prod) => (
                    <TableRow key={prod.id}>
                        <TableCell>{prod.id}</TableCell>
                        <TableCell>{prod.name}</TableCell>
                        <TableCell>{prod.description}</TableCell>
                    </TableRow>
                ))}
            </TableBody>
        </Table>
    </TableContainer> : <Alert severity={"info"}>
        { "No products found" }
    </Alert>)
}

export default ProductsList