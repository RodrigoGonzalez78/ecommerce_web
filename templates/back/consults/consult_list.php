<div>
    <?php if (session()->getFlashdata('success')) {
        echo "
      <div class='mt-3 mb-3 ms-3 me-3 h4 text-center alert alert-success alert-dismissible'>
      <button type='button' class='btn-close' data-bs-dismiss='alert'></button>" . session()->getFlashdata('success') . "
  </div>";
    }
    ?>
</div>
<div class="container m-4">
    <div class="row">
        <div class="col">
            <h4>
                Lista de consultas
            </h4>
        </div>

        <div class="col-auto"> 
            <a href="<?php echo base_url("/archived-list-consults") ?>">
                <button class="btn btn-danger">
                    <img src="assets/img/icons/archive.svg" width="20px" height="20px" alt="Archivar">
                    Archivados
                </button>
            </a>
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
                <th>Atendido/Leido</th>
                <th>Mensaje</th>
                <th>Opciones</th>

            </tr>
        </thead>
        <tbody>
            <?php foreach ($consults as $consult) : ?>
                <tr>
                    <td><?php echo $consult['id']; ?></td>
                    <td><?php echo $consult['name']; ?></td>
                    <td><?php echo $consult['email']; ?></td>
                    <td><?php echo $consult['attended']; ?></td>
                    <td><?php echo $consult['description']; ?></td>
                    <td>
                        <a href="<?php echo base_url('/attended-consults?id='.$consult['id']) ?>">
                            <button class="btn my-btn-primary">
                                <img src="assets/img/icons/done.svg" width="20px" height="20px" alt="Atendido">
                            </button>
                        </a>

                        <a href="<?php echo base_url('/archive-consult?id='.$consult['id']) ?>">
                            <button class="btn btn-danger">
                                <img src="assets/img/icons/archive.svg" width="20px" height="20px" alt="Archivar">
                            </button>
                        </a>

                    </td>

                </tr>
            <?php endforeach; ?>
        </tbody>
    </table>
</div>