<div class="container m-4">
    <div class="row">
        <div class="col">
            <h4>
                Lista de pedidos
            </h4>
        </div>
    </div>
</div>

<div class="container mx-5">

    <table class="table">
        <thead class="table-dark">
            <tr>
                <th>#</th>
                <th>Nombre y Apellido</th>
                <th>Email</th>
                <th>Precio Total</th>
                <th>Fecha</th>
                <th>Factura</th>





            </tr>
        </thead>
        <tbody>
            <?php foreach ($sales as $sale) : ?>
                <tr>
                    <td><?php echo $sale['id']; ?></td>
                    <td><?php echo $sale['name'] . ' ' . $sale['last_name']; ?></td>
                    <td><?php echo $sale['email']; ?></td>
                    <td><?php echo '$' . $sale['total_price']; ?></td>
                    <td><?php echo $sale['date']; ?></td>
                    <td>
                        <a href="<?php echo base_url("/bill?id=" .$sale['id']) ?>">
                        <button class="btn my-btn-primary fw-bold">
                            <img src="assets/img/icons/receipt.svg" width="20px" height="20px" alt="Archivar">
                            Ver
                        </button>
                        </a>
                        <a href="<?php echo base_url("//bill-dowload?id=".$sale['id']) ?>">
                        <button class="btn btn-danger fw-bold">
                            <img src="assets/img/icons/download.svg" width="20px" height="20px" alt="Archivar">
                            Descargar
                        </button>
                        </a>
                       
                    </td>
                </tr>
            <?php endforeach; ?>
        </tbody>
    </table>
</div>