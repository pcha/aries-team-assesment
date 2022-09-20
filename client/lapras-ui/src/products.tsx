import ProductCreator from "./products/create";
import ProductsList from "./products/list";
import {useEffect, useState} from "react";
import ResultNotifier from "./products/resultNotifier";
import {ApiUrl} from "./constants"

function Products(props:{apiURL:string, apiToken:string, logOut:()=>void, filterTerm: string, maxHeight: number}) {
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
        const endpoint = ApiUrl + (props.filterTerm == "" ? "/products" : "/products/search/?q=" + encodeURI(props.filterTerm))
        fetch(endpoint, {
            headers: {
                "Authorization": "Bearer " + props.apiToken
            }
        })
            .then(res => {
                switch (res.status) {
                    case 200:
                        res.json()
                            .then(res => setProducts(res))
                        break
                    case 401:
                        props.logOut()
                        break
                    case 500:
                    default:
                        // TODO Show MUI alert
                        res.json().then((body) => alert(body.error))
                }
            })
    }

    useEffect(fetchProducts, [props.filterTerm])

    const createProduct = (name: string, description: string) => {
        let data = {
            name: name,
            description: description
        }
        fetch(ApiUrl + "/products", {
            method: 'POST',
            headers: {
                Accept: 'application/form-data',
                'Content-Type': 'application/json',
                Authorization: "Bearer " + props.apiToken
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
                        case 401:
                            props.logOut()
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
        <ProductsList products={products} isSearch={props.filterTerm != ""} maxHeight={props.maxHeight} />
    </div>
}

export default Products