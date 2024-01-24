import React, { useState } from 'react';

const AddGoodModal = ({ onClose, refreshGoods }) => {
    const [name, setName] = useState('');
    const [amount, setAmount] = useState(0);

    const handleAddGood = async () => {
        try {
            await fetch('http://localhost:8080/new', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ name, amount: parseInt(amount, 10) }),
            });
            refreshGoods(); 
            onClose(); 
        } catch (error) {
            console.error('Error adding new good:', error);
        }
    };

    const handleSubmit = (e) => {
        e.preventDefault();
        handleAddGood();
    };

    return (
        <div className="modal">
            <form onSubmit={handleSubmit}>
                <label>
                    Name:
                    <input type="text" value={name} onChange={(e) => setName(e.target.value)} />
                </label>
                <label>
                    Amount:
                    <input type="number" value={amount} onChange={(e) => setAmount(e.target.value)} />
                </label>
                <button type="submit">Add</button>
                <button onClick={onClose}>Cancel</button>
            </form>
        </div>
    );
};

export default AddGoodModal;
