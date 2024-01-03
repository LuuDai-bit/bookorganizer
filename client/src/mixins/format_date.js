import dayjs from "dayjs";

const formatDate = (dateString) => {
  if(!dateString) return;

  const date = dayjs(dateString);

  return date.format('MMMM D, YYYY');
}

const formatDateWithFormat = (dateString, format) => {
  if(!dateString) return;

  const date = dayjs(dateString);

  return date.format(format);
}

export default formatDate
export { formatDateWithFormat }
