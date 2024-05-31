const ItemList = ({ items, changeQuantity, deleteItem }) => {
  const changeQuantityHandler = (type, kode) => {
    const typeText = type ? 'Masuk' : 'Keluar';
    const quantity = prompt(`Masukkan jumlah barang ${typeText}`);

    changeQuantity(type, quantity, kode);
  };

  return (
    <>
      <h2>Item List</h2>
      <ul>
        {items.map((item) => (
          <li key={item.kode}>
            {item.kode} - {item.nama} ({item.jumlah}){' '}
            <button onClick={() => changeQuantityHandler(true, item.kode)}>
              Increase
            </button>
            <button onClick={() => changeQuantityHandler(false, item.kode)}>
              Decrease
            </button>
            <button onClick={() => deleteItem(item.kode)}>Delete</button>
          </li>
        ))}
      </ul>
    </>
  );
};

export default ItemList;
