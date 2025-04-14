import React, { useState } from 'react';
import axios from 'axios';

const RequestItem = () => {
    const [item, setItem] = useState({ name: '', shop: '', location: '' });

    const handleChange = (e) => {
        const { name, value } = e.target;
        setItem({ ...item, [name]: value });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            await axios.post('/requests', item);
            alert('Item requested successfully!');
            setItem({ name: '', shop: '', location: '' });
        } catch (error) {
            console.error('Error requesting item:', error);
            alert('Failed to request item');
        }
    };

    return (
        <div className="request-item-form">
            <h2>Request an Item</h2>
            <form onSubmit={handleSubmit}>
                <div className="form-group">
                    <label>Item Name</label>
                    <input
                        type="text"
                        name="name"
                        value={item.name}
                        onChange={handleChange}
                        required
                    />
                </div>
                <div className="form-group">
                    <label>Preferred Shop</label>
                    <input
                        type="text"
                        name="shop"
                        value={item.shop}
                        onChange={handleChange}
                    />
                </div>
                <div className="form-group">
                    <label>Location (optional)</label>
                    <input
                        type="text"
                        name="location"
                        value={item.location}
                        onChange={handleChange}
                    />
                </div>
                <button type="submit" className="submit-btn">Request Item</button>
            </form>
        </div>
    );
};

export default RequestItem;