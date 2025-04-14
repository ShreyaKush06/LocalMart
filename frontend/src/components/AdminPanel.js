import React, { useEffect, useState } from 'react';
import axios from 'axios';

const AdminPanel = () => {
    const [requestedItems, setRequestedItems] = useState([]);

    const fetchRequestedItems = () => {
        axios.get('/requests')
            .then(res => setRequestedItems(res.data))
            .catch(err => console.error(err));
    };

    useEffect(() => {
        fetchRequestedItems();
    }, []);

    const handleListItem = async (item) => {
        try {
            await axios.post('/list-item', item);
            alert('Item listed successfully!');
            fetchRequestedItems(); // Refresh the list
        } catch (error) {
            console.error('Error listing item:', error);
            alert('Failed to list item');
        }
    };

    return (
        <div className="admin-panel">
            <h2>Admin Panel</h2>
            <div className="requested-items">
                {requestedItems.length === 0 ? (
                    <p>No requested items</p>
                ) : (
                    requestedItems.map(item => (
                        <div key={item.id} className="requested-item-card">
                            <h3>{item.name}</h3>
                            <p>Shop: {item.shop}</p>
                            {item.location && <p>📍 {item.location}</p>}
                            <button onClick={() => handleListItem(item)}>List Item</button>
                        </div>
                    ))
                )}
            </div>
        </div>
    );
};

export default AdminPanel;