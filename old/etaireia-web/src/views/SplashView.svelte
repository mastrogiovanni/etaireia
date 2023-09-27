<script>
  import Logo from "../lib/Logo.svelte";
</script>

<div class="logo">
  <Logo />
</div>

<style>
  .logo {
    padding: 20px;
    background: rgb(17, 58, 93);
    box-sizing: border-box;
    font-family: Circular, Arial, sans-serif;
    color: rgb(9, 11, 23);

    -webkit-box-pack: center;
    justify-content: center;
    -webkit-box-align: center;
    align-items: center;
    display: flex;

    /*
            position: absolute;
        
            top: 0px;
            left: 0px;
            width: 100%;
            height: 100%;
        
            */
  }
</style>
