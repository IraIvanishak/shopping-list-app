import React, { useState, useEffect } from 'react';
import AddGoodModal from './AddGoodModal';


const GoodsList = () => {
    const [goods, setGoods] = useState([]);
    const [showModal, setShowModal] = useState(false);
    const [editGood, setEditGood] = useState({ id: null, name: '', amount: 0 });

    const handleEditClick = (good) => {
        setEditGood({ id: good.id, name: good.name, amount: good.amount });
    };

    const handleCancel = () => {
        setEditGood({ id: null, name: '', amount: 0 });
    };

    const handleChange = (e, field) => {
        const value = field === 'amount' ? parseInt(e.target.value, 10) : e.target.value;
        setEditGood({ ...editGood, [field]: value });
    };

    const handleDelete = async (id) => {
        try {
            await fetch(`http://localhost:8080/delete/${id}`, {
                method: 'POST',
            });
            fetchGoods();
        } catch (error) {
            console.error('Error deleting good:', error);
        }
    };

    const handleEdit = async () => {
        try {
            console.log(editGood)
            await fetch(`http://localhost:8080/edit/${editGood.id}`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(editGood),
            });
            fetchGoods();
            setEditGood({ id: null, name: '', amount: 0 });

        } catch (error) {
            console.error('Error updating good:', error);
        }
    };


    // Fetch goods from the API
    const fetchGoods = async () => {
        try {
            const response = await fetch('http://localhost:8080/');
            const data = await response.json();
            setGoods(data);
            console.log(data)
        } catch (error) {
            console.error('Error fetching goods:', error);
        }
    };

    // Fetch goods when the component mounts
    useEffect(() => {
        fetchGoods();
    }, []);

    return (
        <div>
            <ul>
                {goods.map(good => (
                    <li key={good.id}>
                        {editGood.id === good.id ? (
                            <>
                                <input
                                    type="text"
                                    value={editGood.name}
                                    onChange={(e) => handleChange(e, 'name')}
                                />
                                <input
                                    type="number"
                                    value={editGood.amount}
                                    onChange={(e) => handleChange(e, 'amount')}
                                />
                                <button onClick={handleEdit}>Save</button>
                                <button onClick={handleCancel}>Cancel</button>
                            </>
                        ) : (
                            <>
                                {good.name} - Amount: {good.amount}
                                <span onClick={() => handleEditClick(good)} className="icon">
                                    <i className="fas fa-edit"></i>
                                </span>
                                <span onClick={() => handleDelete(good.id)} className="icon">
                                    <i className="fas fa-trash"></i>
                                </span>
                            </>
                        )}
                    </li>
                ))}
            </ul>
            <button onClick={() => setShowModal(true)}>Add New Good</button>

            {showModal && (
                <AddGoodModal onClose={() => setShowModal(false)} refreshGoods={fetchGoods} />
            )}
        </div>
    );
};

export default GoodsList;
