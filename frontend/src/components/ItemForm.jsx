import { useState } from 'react';

const ItemForm = ({ addItem }) => {
  const [kode, setKode] = useState('');
  const [nama, setNama] = useState('');
  const [jumlah, setJumlah] = useState(0);
  const [deskripsi, setDeskripsi] = useState('');
  const [isAktif, setIsAktif] = useState(true);

  const handleSubmit = (e) => {
    e.preventDefault();
    addItem({
      kode,
      nama,
      jumlah,
      deskripsi,
      isAktif,
    });
  };

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <label htmlFor="kode">Kode</label>
        <input
          type="text"
          value={kode}
          id="kode"
          onChange={(e) => setKode(e.target.value)}
        />
      </div>
      <div>
        <label htmlFor="nama">Nama</label>
        <input
          type="text"
          value={nama}
          id="nama"
          onChange={(e) => setNama(e.target.value)}
        />
      </div>
      <div>
        <label htmlFor="jumlah">Jumlah</label>
        <input
          type="number"
          value={jumlah}
          id="jumlah"
          onChange={(e) => setJumlah(e.target.value)}
        />
      </div>
      <div>
        <label htmlFor="deskripsi">Deskripsi</label>
        <textarea
          name="deskripsi"
          value={deskripsi}
          onChange={(e) => setDeskripsi(e.target.value)}
        ></textarea>
      </div>
      <div>
        <label htmlFor="status">Status</label>
        <div>
          <input
            type="radio"
            name="status"
            id="aktif"
            value={true}
            checked={isAktif}
            onChange={() => setIsAktif(true)}
          />
          <label htmlFor="aktif">Aktif</label>
        </div>
        <div>
          <input
            type="radio"
            id="nonaktif"
            value={false}
            checked={!isAktif}
            onChange={() => setIsAktif(false)}
          />
          <label htmlFor="nonaktif">Tidak Aktif</label>
        </div>
      </div>
      <button type="submit">Add Item</button>
    </form>
  );
};

export default ItemForm;
