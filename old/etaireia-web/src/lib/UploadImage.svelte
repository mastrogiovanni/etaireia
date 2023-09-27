<script>
  export let image = undefined;
  export let title = undefined;

  let avatar, fileinput;

  const onFileSelected = async (e) => {
    image = e.target.files[0];
    console.log(typeof image);
    let reader = new FileReader();
    reader.readAsDataURL(image);
    reader.onload = (e) => {
      avatar = e.target.result;
    };
  };

  $: {
    if (!image) {
      avatar = undefined;
    }
  }

</script>

<div id="app">
  <h3>{title}</h3>

  {#if avatar}
    <img class="avatar" src={avatar} alt="d" />
  {:else}

  <div id="imageHolder">
    <img
      
      src="https://cdn4.iconfinder.com/data/icons/small-n-flat/24/user-alt-512.png"
      alt=""
    />
  </div>
  {/if}

  <img
      class="upload"
      src="https://static.thenounproject.com/png/625182-200.png"
      alt=""
      on:click={() => {
        fileinput.click();
      }}/>

  <!--
  <button class="btn btn-sm btn-primary" on:click={() => { fileinput.click() }}>
    Choose Image
  </button>
  -->

  <!--
  <button class="btn btn-sm btn-primary" on:click={submitImage}>
    Invia il documento
  </button>
  -->

  <input
    style="display:none"
    type="file"
    accept=".jpg, .jpeg, .png"
    on:change={(e) => onFileSelected(e)}
    bind:this={fileinput}
  />
</div>

<style>
  #app {
    display: flex;
    align-items: center;
    justify-content: center;
    flex-flow: column;
    border: solid;
  }

  .upload {
    display: flex;
    height: 50px;
    width: 50px;
    cursor: pointer;
  }

  .avatar {
    display: flex;
    height: 200px;
    width: 200px;
  }

  #imageHolder {
    width: 200px;
    height: 200px;
    /*
    line-height: 400px;
    background-color: #cccccc;
    */
  }

  #imageHolder img {
    max-width: 100%;
    max-height: 100%;
    display: inline-block; /* Instead of display: block; */
    margin: 0 auto;
    vertical-align: middle;
  }
</style>
