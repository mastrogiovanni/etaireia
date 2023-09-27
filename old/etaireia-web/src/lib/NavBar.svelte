<script>
    import { page } from "../libs/stores";
    import { PAGE_HOME, PAGE_RECEIVE, PAGE_SIGNED } from "../libs/constants";

    import {
    Collapse,
    Navbar,
    NavbarToggler,
    NavbarBrand,
    Nav,
    NavItem,
    NavLink
  } from 'sveltestrap';

  let isOpen = false;
  const toggle = () => (isOpen = !isOpen);


</script>

<Navbar color="primary" dark >
  <NavbarBrand href="/" class="me-auto">Etaireia</NavbarBrand>
  <NavbarToggler on:click={toggle} class="me-2" />
  <Collapse {isOpen} navbar>
    <Nav navbar>
      <NavItem>
        <NavLink href="#!" on:click={() => { $page = PAGE_HOME; isOpen = false; }}>Home</NavLink>
      </NavItem>
      <NavItem>
        <NavLink href="#!" on:click={() => { $page = PAGE_SIGNED; isOpen = false; }}>Documenti Firmati</NavLink>
      </NavItem>
      <NavItem>
        <NavLink href="#!" on:click={() => { $page = PAGE_RECEIVE; isOpen = false; }}>Ricevi</NavLink>
      </NavItem>
      <NavItem>
        <NavLink href="https://github.com/mastrogiovanni/etaireia">GitHub</NavLink>
      </NavItem>
    </Nav>
  </Collapse>
</Navbar>

<div style="margin-top: 15px;"></div>

<!--
<nav class="navbar navbar-expand-lg navbar-dark bg-primary">
  <div class="container-fluid">
    <a class="navbar-brand" href="#!">Etaireia</a>

    <button
      class="navbar-toggler"
      type="button"
      data-bs-toggle="collapse"
      data-bs-target="#navbarSupportedContent"
      aria-controls="navbarSupportedContent"
      aria-expanded="false"
      aria-label="Toggle navigation"
    >
      <span class="navbar-toggler-icon" /><span class="badge bg-secondary">4</span>
    </button>

    <div class="collapse navbar-collapse" id="navbarSupportedContent">
      <ul class="navbar-nav me-auto mb-2 mb-lg-0">
        <li class="nav-item">
          <a class="nav-link active" aria-current="page" href="#!"  on:click={() => { $page = PAGE_HOME }}>Home</a>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="#!" on:click={() => { $page = PAGE_SIGNED }}>Firmati</a>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="#!" on:click={() => { $page = PAGE_RECEIVE }}>Ricevi</a>
        </li>

        <li class="nav-item dropdown">
          <a
            class="nav-link dropdown-toggle"
            href="#"
            id="navbarDropdown"
            role="button"
            data-bs-toggle="dropdown"
            aria-expanded="false"
          >
            Dropdown
          </a>
          <ul class="dropdown-menu" aria-labelledby="navbarDropdown">
            <li><a class="dropdown-item" href="#">Action</a></li>
            <li><a class="dropdown-item" href="#">Another action</a></li>
            <li><hr class="dropdown-divider" /></li>
            <li><a class="dropdown-item" href="#">Something else here</a></li>
          </ul>
        </li>
        <li class="nav-item">
          <a
            class="nav-link disabled"
            href="#!"
            tabindex="-1"
            aria-disabled="true">Disabled</a
          >
        </li>

      </ul>
      <form class="d-flex">
        <input
          class="form-control me-2"
          type="search"
          placeholder="Search"
          aria-label="Search"
        />
        <button class="btn btn-outline-success" type="submit">Search</button>
      </form>
    </div>

  </div>
</nav>

-->