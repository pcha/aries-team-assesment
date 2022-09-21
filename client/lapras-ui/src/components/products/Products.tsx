import Creator from "./Creator";
import ProductsList from "./List";
import {useState} from "react";
import ResultNotifier from "./ResultNotifier";
import {CreationResult, Product} from "../../hooks/useProducts";

/**
 * Component grouping products related components
 *
 * @param props - {
 *     products: products to list
 *     maxHeight: max height of the product list
 *     handleCreate: handler to call to create a product
 * }
 * @constructor
 */
function Products(props: { products: Product[], maxHeight: number, handleCreate: (name: string, description: string) => Promise<CreationResult> }) {
    const [showResult, setShowResult] = useState(false)
    const [resultMessage, setResultMessage] = useState("")
    const [success, setSuccess] = useState(false)

    const showResultDialog = (message: string, success: boolean) => {
        setResultMessage(message)
        setSuccess(success)
        setShowResult(true)
    }

    return <div>
        <ResultNotifier open={showResult} setOpen={setShowResult} message={resultMessage} success={success}/>
        <Creator
            handleCreate={(name, description) => props.handleCreate(name, description).then((res) => showResultDialog(res.message, res.success))}/>
        <ProductsList products={props.products} maxHeight={props.maxHeight}/>
    </div>
}

export default Products