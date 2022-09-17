import './App.css'
import ProductsList from "./products/list";
import ProductCreator from "./products/create"
import Products from "./products/container";

function App() {

  return (
    <div className="App">
        <h2>Lapras</h2>
        <Products apiURL={import.meta.env.API_URL} />
    </div>
  )
}

export default App
