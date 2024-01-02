import dayjs from "dayjs";

const formatDate = (dateString) => {
  const date = dayjs(dateString);

  return date.format('MMMM D, YYYY');
}

export default formatDate
