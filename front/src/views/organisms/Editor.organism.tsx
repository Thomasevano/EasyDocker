import React, { type ChangeEvent, useState } from 'react'
import useEditor from '../../hooks/useEditor'
import { type TDrawer } from '../../types/Drawer'
import EventsCanvas from '../../services/canvas/Events.canvas'
import Input from '../atoms/Forms/Input.atom'
import Button from '../atoms/Forms/Button.atom'

const EditorOrganism = ({ drawer }: { drawer: TDrawer }): JSX.Element => {
  const { fields } = useEditor(drawer)

  const [formDrawer, setFormDrawer] = useState<TDrawer>(drawer)

  const changeValue: (event: ChangeEvent<HTMLInputElement>)
  => void =
    (event: ChangeEvent<HTMLInputElement>): void => {
      setFormDrawer({ ...formDrawer, [event.target.name]: event.target.value })
    }

  const onClose = (): void => {
    EventsCanvas.clearSelectedDrawer()
    EventsCanvas.updateScreen()
  }

  const onSubmit = (): void => {

  }

  const submitText: string = drawer?.entity?.id !== null ? 'Update' : 'Create'

  // yup validation
  // const validators: object = fields.reduce((acc, field) => ({ [field.name]: field.validator }), {})

  return (
    <div className="w-full h-full border-l-2 ">
      <div className="border-b-2 p-2 flex flex items-center justify-between">
        <h2><strong>Editor</strong></h2>

        <Button
          className="bg-transparent text-blue-500 hover:text-white font-bold h-2"
          label={'X'}
          onClick={onClose}
        />
      </div>

      <form className="p-2" onSubmit={onSubmit}>

        {fields.map((field, index) => (
          <Input
            label={field.label}
            type={field.type}
            name={field.name}
            onChange={changeValue}
            key={index}
          />
        ))}

        <div className="mt-5">
          <Button
            label={submitText}
            onClick={onSubmit}
            className={'w-full'}
            direction={'right'}
          />
        </div>

      </form>
    </div>
  )
}

export default EditorOrganism