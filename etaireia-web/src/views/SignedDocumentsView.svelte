<script>
  import { Table } from "sveltestrap";
  import moment from "moment";
  import "moment/locale/it";
  moment.locale("it");

  import {
    Card,
    CardBody,
    CardFooter,
    CardHeader,
    CardSubtitle,
    CardText,
    CardTitle,
  } from "sveltestrap";
  import { onMount } from "svelte";

  onMount(async () => {
    const response = await fetch("/api/v1/signature", {
      method: "POST",
      body: formData,
    });
  }) 

</script>

<Card class="mb-3">
  <CardHeader>
    <CardTitle>Documenti Firmati</CardTitle>
  </CardHeader>
  <CardBody>
    <Table hover>
      <thead>
        <tr>
          <th>#</th>
          <th>Documento</th>
          <th>Data</th>
          <th />
        </tr>
      </thead>
      <tbody>
        <tr>
          <th scope="row">1</th>
          <td>Segreteria</td>
          <td>{moment().format("MMM Do YY")}</td>
          <td><button class="btn"><i class="bi bi-cloud-download" /></button></td>
        </tr>
        <tr>
          <th scope="row">2</th>
          <td>Prof. Xevier</td>
          <td>{moment().format("MMM Do YY")}</td>
          <td><button class="btn"><i class="bi bi-cloud-download" /></td>
        </tr>
        <tr>
          <th scope="row">3</th>
          <td>...</td>
          <td>{moment().format("MMM Do YY")}</td>
          <td><button class="btn"><i class="bi bi-cloud-download" /></td>
        </tr>
      </tbody>
    </Table>
  </CardBody>
  <CardFooter />
</Card>
