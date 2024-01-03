import dayjs from "dayjs";

const formatDate = (dateString) => {
  if(!dateString) return;

  const date = dayjs(dateString);

  return date.format('MMMM D, YYYY');
}

export default formatDate
