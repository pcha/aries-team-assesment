import ProductCreator from "./create";
import ProductsList from "./list";
import {useEffect, useState} from "react";
import ResultNotifier from "./resultNotifier";

function Products(props:{apiURL:string}) {
    const [products, setProducts] = useState<any[]>([])
    const [showResult, setShowResult] = useState(false)
    const [resultMessage, setResultMessage] = useState("")
    const [success, setSuccess] = useState(false)

    const showResultDialog = (message: string, success: boolean) => {
        setResultMessage(message)
        setSuccess(success)
        setShowResult(true)
    }

    const fetchProducts = () => {
        fetch(import.meta.env.VITE_API_URL + "/products").
        then(res => res.json())
            .then(res => setProducts(res))
    }
    useEffect(fetchProducts, [])

    const createProduct = (name: string, description: string) => {
        let data = {
            name: name,
            description: description
        }
        fetch(import.meta.env.VITE_API_URL + "/products", {
            method: 'POST',
            headers: {
                Accept: 'application/form-data',
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data),
        })
            .then(
                (res) => {
                    switch (res.status) {
                        case 201:
                            showResultDialog("Product created.", true)
                            break
                        case 400:
                            showResultDialog("Invalid data.", false)
                            break
                        case 500:
                        default:
                            showResultDialog("Try again later.", false)

                    }
                    setShowResult(true)
                },
                (err) => {
                    alert(err)
                    console.log(err)
                }
            ).then(() => fetchProducts())
    }
    return <div>
        <ResultNotifier open={showResult} setOpen={setShowResult} message={resultMessage} success={success} />
        <ProductCreator create={createProduct}/>
        <ProductsList products={products} />
    </div>
}

export default Products