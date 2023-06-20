import { NextApiRequest, NextApiResponse } from 'next'
import { people } from '../../../data'
import type { Person, ResponseError } from '../../../interfaces'
import tracker from '@middleware.io/agent-apm-nextjs';

export default function personHandler(
  req: NextApiRequest,
  res: NextApiResponse<Person | ResponseError>
) {
  const { query } = req
  const { id } = query
  const person = people.find((p) => p.id === id)

  if (person) {

    tracker.info(`Person with id ${id} found successfully.`, person);
    return res.status(200).json(person)

  } else {

    tracker.error(`Requested person with id ${id} not found.`, { id: id });
    return res.status(404).json({ message: `User not found.` })

  }
}
