const ItemList = ({ items }) => {
  return (
    <>
      <h2>Item List</h2>
      <ul>
        {items.map((item) => (
          <li key={item.kode}>
            {item.kode} - {item.nama} ({item.jumlah})
          </li>
        ))}
      </ul>
    </>
  );
};

export default ItemList;
