﻿using System.ComponentModel.DataAnnotations;

namespace Education.Domain
{
    public class Curso
    {
        [Key]
        public Guid CursoId { get; set; }

        [Required]
        [StringLength(100)]
        public string Titulo { get; set; }

        [Required]
        [StringLength(1000)]
        public string Descripcion { get; set; }

        [DataType(DataType.Date)]

        public DateTime? FechaPublicacion { get; set; }

        [DataType(DataType.Date)]
        public DateTime? FechaCreacion { get; set; }

        public decimal Precio { get; set; }
    }
}