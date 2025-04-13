import React, { useState } from 'react';
import axios from 'axios';

const AddProduct = ({ onProductAdded }) => {
    const [product, setProduct] = useState({
        name: '',
        price: '',
        shop: '',
        onBlinkit: false,
        location: ''
    });

    const handleChange = (e) => {
        const { name, value, type, checked } = e.target;
        setProduct({
            ...product,
            [name]: type === 'checkbox' ? checked : value
        });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        
        // Generate a random ID since we don't have proper DB
        const newProduct = {
            ...product,
            id: Math.floor(Math.random() * 1000)
        };

        try {
            await axios.post('/products', newProduct);
            setProduct({ name: '', price: '', shop: '', onBlinkit: false, location: '' });
            if (onProductAdded) onProductAdded();
            alert('Product added successfully!');
        } catch (error) {
            console.error('Error adding product:', error);
            alert('Failed to add product');
        }
    };

    return (
        <div className="add-product-form">
            <h2>Add New Product</h2>
            <form onSubmit={handleSubmit}>
                <div className="form-group">
                    <label>Product Name</label>
                    <input
                        type="text"
                        name="name"
                        value={product.name}
                        onChange={handleChange}
                        required
                    />
                </div>
                
                <div className="form-group">
                    <label>Price (with ₹ symbol)</label>
                    <input
                        type="text"
                        name="price"
                        value={product.price}
                        onChange={handleChange}
                        placeholder="₹150"
                        required
                    />
                </div>
                
                <div className="form-group">
                    <label>Shop Name</label>
                    <input
                        type="text"
                        name="shop"
                        value={product.shop}
                        onChange={handleChange}
                        required
                    />
                </div>
                
                <div className="form-group checkbox">
                    <label>
                        <input
                            type="checkbox"
                            name="onBlinkit"
                            checked={product.onBlinkit}
                            onChange={handleChange}
                        />
                        Available on Blinkit
                    </label>
                </div>
                
                <div className="form-group">
                    <label>Location (optional)</label>
                    <input
                        type="text"
                        name="location"
                        value={product.location}
                        onChange={handleChange}
                        placeholder="Building/Floor/Room"
                    />
                </div>
                
                <button type="submit" className="submit-btn">Add Product</button>
            </form>
        </div>
    );
};

export default AddProduct;