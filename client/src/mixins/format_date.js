import dayjs from "dayjs";

const formatDate = (dateString) => {
  const date = dayjs(dateString);

  if(dateString == '0001-01-01T00:00:00Z') return null;

  return date.format('MMMM D, YYYY');
}

export default formatDate
