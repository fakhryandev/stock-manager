import { useEffect, useState } from 'react';
import ItemList from './components/ItemList';
import ItemForm from './components/ItemForm';
import axios from 'axios';

function App() {
  const [items, setItems] = useState([]);

  useEffect(() => {
    getItems();
  }, []);

  const getItems = async () => {
    const response = await axios.get(`http://localhost:8080/api/item`);
    const data = response.data.data;
    setItems(data);
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
    } catch (error) {
      console.log(error);
      console.log(error);
      // alert(error);
    }
  };

  const changeQuantity = async (type, quantity, kode) => {
    const typeRequest = type ? 'increase' : 'decrease';
    const body = {
      jumlah: parseInt(quantity),
    };
    const response = await axios.patch(
      `http://localhost:8080/api/item/${kode}/${typeRequest}`,
      body
    );

    console.log(response);
  };

  const deleteItem = async (kode) => {
    const response = await axios.delete(
      `http://localhost:8080/api/item/${kode}`
    );

    console.log(response);
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
      />
    </>
  );
}

export default App;
