import React, { useEffect, useState } from 'react';
import axios from 'axios';
import AddProduct from './AddProduct';

const ProductList = () => {
    const [products, setProducts] = useState([]);
    const [filterNonBlinkit, setFilter] = useState(true);
    const [showAddForm, setShowAddForm] = useState(false);

    const fetchProducts = () => {
        axios.get('/products')
            .then(res => setProducts(res.data))
            .catch(err => console.error(err));
    };

    useEffect(() => {
        fetchProducts();
    }, []);

    const handleProductAdded = () => {
        fetchProducts(); // Refresh the list
        setShowAddForm(false); // Hide the form
    };

    return (
        <div className="product-container">
            <div className="product-header">
                <div className="filter-section">
                    <label>
                        <input
                            type="checkbox"
                            checked={filterNonBlinkit}
                            onChange={() => setFilter(!filterNonBlinkit)}
                        />
                        Show Only Non-Blinkit Items
                    </label>
                </div>
                <button 
                    className="add-product-btn"
                    onClick={() => setShowAddForm(!showAddForm)}
                >
                    {showAddForm ? "Cancel" : "Add Product"}
                </button>
            </div>
            
            {showAddForm && <AddProduct onProductAdded={handleProductAdded} />}
            
            <div className="product-grid">
                {products.length === 0 ? (
                    <p>No products found</p>
                ) : products
                    .filter(product => filterNonBlinkit ? !product.onBlinkit : true)
                    .map(product => (
                        <div key={product.id} className="product-card">
                            <h3>{product.name}</h3>
                            <p>Price: {product.price}</p>
                            <p>Shop: {product.shop}</p>
                            {product.location && <p>📍 {product.location}</p>}
                        </div>
                    ))}
            </div>
        </div>
    );
};

export default ProductList;