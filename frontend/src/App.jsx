import { useState } from 'react';
import ItemList from './components/ItemList';

function App() {
  const [items, setItems] = useState([]);

  return (
    <>
      <ItemList items={items} />
    </>
  );
}

export default App;
