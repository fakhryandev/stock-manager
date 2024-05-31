import { useEffect, useState } from 'react';
import ItemList from './components/ItemList';
import ItemForm from './components/ItemForm';
import axios from 'axios';

function App() {
  const [items, setItems] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    getItems();
  }, []);

  const getItems = async () => {
    try {
      const response = await axios.get(`http://localhost:8080/api/item`);
      const data = response.data.data;
      setItems(data);
    } finally {
      setLoading(false);
    }
  };

  const addItem = async (itemData) => {
    try {
      const { isAktif, jumlah } = itemData;
      const body = {
        ...itemData,
        status: isAktif,
        jumlah: parseInt(jumlah),
      };
      const response = await axios.post(`http://localhost:8080/api/item`, body);
      setItems([...items, body]);
      const message = response.data.message;
      alert(message);
    } catch (error) {
      console.log(error);
      const errMessage = error.response.data.message;
      alert(errMessage);
    }
  };

  const changeQuantity = async (type, quantity, kode) => {
    try {
      const typeRequest = type ? 'increase' : 'decrease';
      const body = {
        jumlah: parseInt(quantity),
      };
      const response = await axios.patch(
        `http://localhost:8080/api/item/${kode}/${typeRequest}`,
        body
      );

      console.log(response);

      const updatedItem = response.data.data;
      setItems((prevItems) =>
        prevItems.map((item) => (item.kode === kode ? updatedItem : item))
      );
    } catch (error) {
      console.log(error);
      const message = error.response.data.message;
      alert(message);
    }
  };

  const deleteItem = async (kode) => {
    try {
      const response = await axios.delete(
        `http://localhost:8080/api/item/${kode}`
      );

      setItems((prevItems) => prevItems.filter((item) => item.kode !== kode));
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <>
      <header className="p-4 bg-slate-900">
        <h1 className="text-3xl text-white">Stock Manager</h1>
      </header>
      <ItemForm addItem={addItem} />
      <ItemList
        items={items}
        changeQuantity={changeQuantity}
        deleteItem={deleteItem}
        isLoading={loading}
      />
    </>
  );
}

export default App;
