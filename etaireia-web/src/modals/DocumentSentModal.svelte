
<script>
    import { Button, Modal } from "sveltestrap";
    import { credential } from '../libs/stores';

    export let isOpen = false;

    function toggle() {
        isOpen = !isOpen
    }
</script>

<Modal body {isOpen} {toggle} header="Inviato!!!">
    <p>
        Documento inviato firmato con successo.
    </p>
    <img
      src="https://static6.depositphotos.com/1110975/606/i/950/depositphotos_6063518-stock-photo-signed-document.jpg"
      alt="Documento Inviato"
      class="img-fluid"
    />
  </Modal>
  