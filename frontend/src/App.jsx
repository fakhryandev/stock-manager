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
    const { isAktif, jumlah } = itemData;
    const body = {
      ...itemData,
      status: isAktif,
      jumlah: parseInt(jumlah),
    };
    const response = await axios.post(`http://localhost:8080/api/item`, body);
    const { status } = response.data;
  };

  return (
    <>
      <ItemForm addItem={addItem} />
      <ItemList items={items} />
    </>
  );
}

export default App;
