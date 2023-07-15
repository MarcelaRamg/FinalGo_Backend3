// function convertDateFormat(dateString) {
//     // Eliminar el sufijo ordinal del día
//     const day = dateString.replace(/(st|nd|rd|th)/g, '');
  
//     const inputFormat = 'D MMM YYYY';
//     const outputFormat = 'YYYY-MM-DD';
  
//     const inputDate = moment(day, inputFormat);
//     const outputDate = inputDate.format(outputFormat);
  
//     return outputDate;
//   }
  
//   const date1 = '1st March 1974';
//   const date2 = '22nd January 2013';
//   const date3 = '7th April 1904';
  
//   const convertedDate1 = convertDateFormat(date1);
//   const convertedDate2 = convertDateFormat(date2);
//   const convertedDate3 = convertDateFormat(date3);
  
//   console.log(convertedDate1);
//   console.log(convertedDate2);
//   console.log(convertedDate3);

  function convertirFecha(fecha) {
    // Separar la cadena de fecha en día, mes y año
    var partes = fecha.split(' ');
    var dia = partes[0];
    var mes = partes[1];
    var anio = partes[2];
  
    // Mapear el nombre del mes a su número correspondiente
    var meses = {
      "January": "01", "February": "02", "March": "03", "April": "04", "May": "05", "June": "06",
      "July": "07", "August": "08", "Septenmber": "09", "October": "10", "November": "11", "December": "12"
    };
  
    // Obtener el número del mes correspondiente
    var numeroMes = meses[mes];
  
    // Formatear la fecha en el formato "AAAA-MM-DD"
    var fechaFormateada = anio + '-' + numeroMes + '-' + dia;
  
    return fechaFormateada;
  }
  
  // Ejemplos de uso
  var fecha1 = "1st March 1974";
  var fecha2 = "22nd January 2013";
  var fecha3 = "7th April 1904";
  
  console.log(convertirFecha(fecha1)); // Resultado: "1974-03-01"
  console.log(convertirFecha(fecha2)); // Resultado: "2013-01-22"
  console.log(convertirFecha(fecha3)); // Resultado: "1904-04-07"